pipeline{
    agent any

    stages{

        stage('Test'){
             steps{
                 sh 'make tests'
             }
        }
        stage('Build'){
                    steps{
                        sh 'make build'
                    }
                }
    }



}