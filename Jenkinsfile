pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        git(url: 'https://github.com/nengbai/study.git', branch: 'master')
      }
    }

    stage('test') {
      steps {
        sh 'ls -ltr'
      }
    }

    stage('Deliver') {
      steps {
        sh 'pwd'
      }
    }

    stage('input') {
      steps {
        echo 'welcome to JD!'
        echo 'Congratulation! You success to pass the test!'
      }
    }

  }
  environment {
    CI = 'True'
  }
}