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

        stage('Build') {
            steps {
                bat 'go build -o prog.exe prog.go'
            }
        }
        stage('Unit Test') {
            steps {
                bat 'go test ./...'
            }
        }

        stage('Run') {
            steps {
                bat '.\\prog.exe'
            }
        }
    }
}

