package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Product struct to match your API
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

func main() {
	baseURL := "http://localhost:9090"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Println("🚀 Starting API Test Client...")
	fmt.Println(strings.Repeat("=", 50))

	// Test 1: Default endpoint (/)
	fmt.Println("\n1️⃣ Testing Default Endpoint (/)")
	testDefaultEndpoint(client, baseURL)

	// Test 2: Hello World endpoint (/helloworld)
	fmt.Println("\n2️⃣ Testing Hello World Endpoint (/helloworld)")
	testHelloWorldEndpoint(client, baseURL)

	// Test 3: Get all products
	fmt.Println("\n3️⃣ Testing GET Products Endpoint (/products)")
	testGetProducts(client, baseURL)

	// Test 4: Add a new product
	fmt.Println("\n4️⃣ Testing POST Products Endpoint (/products)")
	testAddProduct(client, baseURL)

	// Test 5: Update an existing product
	fmt.Println("\n5️⃣ Testing PUT Products Endpoint (/products)")
	testUpdateProduct(client, baseURL)

	fmt.Println("\n✅ All tests completed!")
}

func testDefaultEndpoint(client *http.Client, baseURL string) {
	data := "Hello from Go Test Client"
	req, err := http.NewRequest("POST", baseURL+"/", bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}

	fmt.Printf("📤 Request: POST %s/ with data: %s\n", baseURL, data)
	fmt.Printf("📥 Response Status: %s\n", resp.Status)
	fmt.Printf("📥 Response Body: %s\n", string(body))
}

func testHelloWorldEndpoint(client *http.Client, baseURL string) {
	data := "Amar"
	req, err := http.NewRequest("POST", baseURL+"/helloworld", bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}

	fmt.Printf("📤 Request: POST %s/helloworld with data: %s\n", baseURL, data)
	fmt.Printf("📥 Response Status: %s\n", resp.Status)
	fmt.Printf("📥 Response Body: %s\n", string(body))
}

func testGetProducts(client *http.Client, baseURL string) {
	req, err := http.NewRequest("GET", baseURL+"/products", nil)
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}

	fmt.Printf("📤 Request: GET %s/products\n", baseURL)
	fmt.Printf("📥 Response Status: %s\n", resp.Status)
	fmt.Printf("📥 Response Body: %s\n", string(body))
}

func testAddProduct(client *http.Client, baseURL string) {
	newProduct := Product{
		ID:          3,
		Name:        "Cappuccino",
		Description: "Italian Coffee with Foam",
		Price:       11.99,
		SKU:         "cap001",
	}

	jsonData, err := json.Marshal(newProduct)
	if err != nil {
		fmt.Printf("❌ Error marshaling JSON: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", baseURL+"/products", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}

	fmt.Printf("📤 Request: POST %s/products with data: %s\n", baseURL, string(jsonData))
	fmt.Printf("📥 Response Status: %s\n", resp.Status)
	fmt.Printf("📥 Response Body: %s\n", string(body))
}

func testUpdateProduct(client *http.Client, baseURL string) {
	updatedProduct := Product{
		ID:          1,
		Name:        "Espresso Updated",
		Description: "Strong Bitter Coffee - Updated",
		Price:       10.99,
		SKU:         "esp001",
	}

	jsonData, err := json.Marshal(updatedProduct)
	if err != nil {
		fmt.Printf("❌ Error marshaling JSON: %v\n", err)
		return
	}

	req, err := http.NewRequest("PUT", baseURL+"/products", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}

	fmt.Printf("📤 Request: PUT %s/products with data: %s\n", baseURL, string(jsonData))
	fmt.Printf("📥 Response Status: %s\n", resp.Status)
	fmt.Printf("📥 Response Body: %s\n", string(body))
}
