pipeline{
    agent any

    stages{

        stage('Test'){
             steps{
                 sh 'make tests'
             }
        }
        stage('Sonar'){
                    steps{
                        sh 'make sonar-scanner
'
                    }
                }
        stage('Build'){
                    steps{
                        sh 'make build'
                    }
                }
    }


}
