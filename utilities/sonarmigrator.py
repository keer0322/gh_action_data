import requests

# Set source SonarCloud organization and access token
source_org = "source_organization"
source_token = "source_access_token"

# Set target SonarCloud organization and access token
target_org = "target_organization"
target_token = "target_access_token"

# Number of items to fetch per page
page_size = 100

# API endpoint to fetch all project keys from source instance
source_project_keys_api = f"https://source.sonarcloud.io/api/components/search?qualifiers=TRK&organization={source_org}&ps={page_size}"

# API endpoint to create project on target instance
target_create_project_api = "https://target.sonarcloud.io/api/projects/create"

# Initialize lists to store successful and failed projects
successful_projects = []
failed_projects = []

# Fetch project keys from source instance
response = requests.get(source_project_keys_api, auth=(source_token,))

if response.status_code == 200:
    source_projects = response.json().get("components", [])
    for project in source_projects:
        project_key = project["key"]
        project_name = project["name"]
        project_language = project["language"]

        # Create project on target instance
        payload = {
            "name": project_name,
            "project": project_key,
            "language": project_language,
            "organization": target_org
        }
        response = requests.post(target_create_project_api, json=payload, auth=(target_token,))

        if response.status_code == 200:
            successful_projects.append(project_key)
            print(f"Project {project_key} migrated to target instance with response: {response.text}")
        else:
            failed_projects.append(project_key)
            print(f"Failed to migrate project {project_key} to target instance with response: {response.text}")

    # Print list of successful projects
    print("Successful projects:")
    for project in successful_projects:
        print(project)

    # Print list of failed projects
    print("Failed projects:")
    for project in failed_projects:
        print(project)
else:
    print(f"Failed to fetch project keys from source instance with status code: {response.status_code}")
