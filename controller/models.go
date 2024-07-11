package controller

import (
	"encoding/json"
	"fmt"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
	"io/ioutil"
	"net/http"
)

func ChatPredictions(w http.ResponseWriter, r *http.Request) {
	// Load the ONNX model file
	modelData, err := ioutil.ReadFile("path/to/your/model.onnx")
	if err != nil {
		http.Error(w, "Failed to load model file", http.StatusInternalServerError)
		return
	}

	// Initialize the Gorgonnx backend
	backend := gorgonnx.NewGraph()

	// Initialize the ONNX model with the Gorgonnx backend
	model := onnx.NewModel(backend)

	// Unmarshal the model
	err = model.UnmarshalBinary(modelData)
	if err != nil {
		http.Error(w, "Failed to unmarshal model", http.StatusInternalServerError)
		return
	}

	// Read input question from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	var requestData map[string]string
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	question := requestData["question"]

	// Preprocess the question and create the input tensor for the model
	// Adjust preprocessing according to your model's requirements
	inputShape := []int{1, len(question)} // Assuming your input shape; adjust as needed
	inputTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(inputShape...))
	inputTensorData := inputTensor.Data().([]float32)

	// Fill inputTensorData with your input question data
	for i, char := range question {
		inputTensorData[i] = float32(char)
	}

	// Set the input tensor in the model
	err = model.SetInput(0, inputTensor)
	if err != nil {
		http.Error(w, "Failed to set input tensor", http.StatusInternalServerError)
		return
	}

	// Run inference using the Gorgonnx backend
	err = backend.Run()
	if err != nil {
		http.Error(w, "Failed to run inference", http.StatusInternalServerError)
		return
	}

	// Get the output tensor
	outputTensor, err := model.GetOutputTensors()
	if err != nil {
		http.Error(w, "Failed to get output tensor", http.StatusInternalServerError)
		return
	}

	output := fmt.Sprintf("Output: %v", outputTensor)

	// Send the prediction result as the response
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"prediction": output}
	json.NewEncoder(w).Encode(response)
}
