#!/usr/bin/env groovy
pipeline{
    agent any

    stages{

        stage('Test'){
             steps{
                 echo 'make'
             }
        }
        stage('Sonar'){
                    steps{
                        echo 'make sonar-scanner'
                    }
                }
        stage('Build'){
                    steps{
                        echo 'make build'
                    }
                }
    }




}
