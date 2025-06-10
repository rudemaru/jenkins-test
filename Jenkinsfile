pipeline {
    agent {
      node {
        label 'docker-agent-alpine'
      }
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh'''
                go build main.go
                echo 'Build is successful'
                '''
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                sh'''
                sleep 5
                echo 'Tests passed successfully'
                '''
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                sh'''
                sleep 5
                echo 'Deployed successfully'
                '''
            }
        }
    }
}
