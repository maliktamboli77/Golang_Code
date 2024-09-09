/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'php:8.3.9-alpine3.20' } }
    stages {
        stage('build') {
            steps {
                sh 'php --version'
            }
        }
    }
}
