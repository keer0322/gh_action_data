package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Project struct {
    Key string `json:"key"`
}

type ProjectSearchResponse struct {
    Components []Project `json:"components"`
    Paging     Paging    `json:"paging"`
}

type Paging struct {
    NextPage string `json:"nextPage"`
}

func main() {
    // Set your SonarCloud organization
    organization := "your_organization"

    // Set your SonarCloud access token
    token := "your_access_token"

    // Number of items to fetch per page
    pageSize := 100

    // API endpoint to fetch all project keys
    projectKeysAPI := fmt.Sprintf("https://sonarcloud.io/api/components/search?qualifiers=TRK&organization=%s&ps=%d", organization, pageSize)

    // Create HTTP client
    client := &http.Client{}

    // Fetch project keys and information using pagination
    for {
        // Create HTTP request
        req, err := http.NewRequest("GET", projectKeysAPI, nil)
        if err != nil {
            fmt.Println("Error creating request:", err)
            return
        }

        // Set Authorization header
        req.SetBasicAuth(token, "")

        // Send request
        resp, err := client.Do(req)
        if err != nil {
            fmt.Println("Error sending request:", err)
            return
        }
        defer resp.Body.Close()

        // Read response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            return
        }

        // Parse JSON response
        var projects ProjectSearchResponse
        if err := json.Unmarshal(body, &projects); err != nil {
            fmt.Println("Error parsing JSON:", err)
            return
        }

        // Loop through project keys and fetch project information
        for _, project := range projects.Components {
            // API endpoint to fetch project details
            apiEndpoint := fmt.Sprintf("https://sonarcloud.io/api/components/show?key=%s", project.Key)
            fmt.Println("Fetching information for project with key:", project.Key)

            // Create HTTP request
            req, err := http.NewRequest("GET", apiEndpoint, nil)
            if err != nil {
                fmt.Println("Error creating request:", err)
                return
            }

            // Set Authorization header
            req.SetBasicAuth(token, "")

            // Send request
            resp, err := client.Do(req)
            if err != nil {
                fmt.Println("Error sending request:", err)
                return
            }
            defer resp.Body.Close()

            // Read response body
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                fmt.Println("Error reading response body:", err)
                return
            }

            // Print response body
            fmt.Println(string(body))
            fmt.Println() // Add newline for readability
        }

        // Check if there are more pages
        if projects.Paging.NextPage == "" {
            break
        }

        // Update API endpoint for next page
        projectKeysAPI = "https://sonarcloud.io" + projects.Paging.NextPage
    }
}
