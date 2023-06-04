pipeline {
    agent any
 

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/iorgu12/go1.git'
            }
        }
        
        stage('Test') {
            steps {
                bat 'go test'
            }
        }
        

        stage('Build') {
            steps {
                bat 'go build -o prog.exe prog.go'
            }
        }
        

         stage('Deploy') {
    steps {
        echo 'Deploying....'
        bat '''
             scp -i SSH_KEY prog.exe coco@192.168.81.129:/home/iorgu/lab
             ssh -i SSH_KEY coco@192.168.81.129 "chmod +x /home/iorgu/lab/prog.exe && /home/iorgu/lab/prog.exe"
        '''
    }
}

    }
}

