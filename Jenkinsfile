#!/usr/bin/env groovy
pipeline{
    agent any

    stages{

        stage('Test'){
             steps{
                 sh 'make'
             }
        }
        stage('Sonar'){
                    steps{
                        sh 'make sonar-scanner'
                    }
                }
        stage('Build'){
                    steps{
                        sh 'make build'
                    }
                }
    }


}
