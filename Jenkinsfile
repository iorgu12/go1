pipeline {
    agent any
    environment {
        // Set up the GOPATH
        GOPATH = "${WORKSPACE}"
        GOOS = "windows"
        GOARCH = "amd64"
    }
    tools {
        // Make sure 'go-1.20.4' is installed on Jenkins
        go 'go-1.20.4'
    }
    stages {
        stage('Checkout') {
            steps {
                // Fetch source code from your repository
                git 'https://github.com/iorgu12/go.git'
            }
        }
        stage('Build') {
            steps {
                // Build the source code to produce a binary artifact
                bat 'go build -o prog.go'
            }
        }
    }
}
