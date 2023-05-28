pipeline {
    agent any
    environment {
        // Set up the GOPATH
        GOPATH = "${WORKSPACE}"
        GOOS = "windows"
        GOARCH = "amd64"
        
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/iorgu12/go.git'
            }
        }

        stage('Initialize Go module') {
            steps {
                bat 'go mod init github.com/iorgu12/go'  // Replace with your actual github username and repository name
            }
        }

        stage('Unit Test') {
            steps {
                bat 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                bat 'go build -o prog.exe prog.go'
            }
        }

        stage('Run') {
            steps {
                bat '.\\prog.exe'
            }
        }
    }
}
