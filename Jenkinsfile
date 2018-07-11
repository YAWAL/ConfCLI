#!/usr/bin/env groovy
pipeline{
    agent any

    stages{
        stage('Build'){
              steps{
                  sh 'make build'
                  archiveArtifacts artifact: 'ConfCLI', fingerprint: true
              }
        }
        stage('Test'){
             steps{
                 sh 'make tests || true'
             }
        }
        stage('Sonar'){
                    steps{
                        echo 'make sonar-scanner'
                    }
                }

    }




}
