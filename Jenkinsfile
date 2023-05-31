pipeline {
    agent any
 

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/iorgu12/go.git'
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

        stage('Run') {
            steps {
                bat '.\\prog.exe'
            }
        }
    }
    stage('Deploy') {
      steps {
        withCredentials([sshUserPrivateKey(credentialsId: 'your-ssh-credential-id', keyFileVariable: 'SSH_KEY')]) {
          script {
            // Copy the binary to the target VM
            sshagent(['your-ssh-credential-id']) {
              sh "ssh -o StrictHostKeyChecking=no -i $SSH_KEY username@iorgu 'mkdir -p /path/to/destination && scp -i $SSH_KEY C:\ProgramData\Jenkins\.jenkins\workspace\go>.\prog.exe username@iorgu-vm:/path/to/destination'"
            }
            
            // Start the program on the target VM
            sshagent(['your-ssh-credential-id']) {
              sh "ssh -o StrictHostKeyChecking=no -i $SSH_KEY username@iorgu 'nohup /path/to/destination/binary > /dev/null 2>&1 &'"
   
            }
          }
        }
      }
    }
}
