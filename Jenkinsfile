#!/usr/bin/env groovy
pipeline{
    agent any

    stages{
        stage('Build'){
              steps{
                  sh 'apd-get install golang-go'
                  sh 'make install-helpers'
                  sh 'make build || true'
                  archiveArtifacts artifacts: 'ConfCLI', fingerprint: true
              }
        }
        stage('Test'){
             steps{
                 sh 'make tests || true'
             }
        }
        stage('Sonar'){
                    steps{
                        'make sonar-scanner'
                    }
        }
        post {
            failure{
            mail to: yavorskyyval@gmail.com, subject: 'jenkins'
            }
        }

    }




}
