pipeline {
    agent any
    triggers {
        githubPush()
    }

    environment {
        IMAGE_NAME = "saimanasg/webapp-hello-world"
        DOCKER_CREDENTIALS_ID = "CloudJenkinsDockerHubPAT"
        TAG = "${env.VERSION}" // Will be set in the 'Get Version' stage
    }

    stages {
        stage('Set up Buildx') {
            steps {
                script {
                    sh 'docker run --rm --privileged multiarch/qemu-user-static --reset -p yes'
                    sh 'docker buildx create --use --name multiarch_builder || true'
                }
            }
        }

        stage('Get Version') {
            steps {
                script {
                    // Get the latest git tag
                    def currentVersion = sh(script: 'git describe --tags --abbrev=0', returnStdout: true).trim()
                    echo "Current version from Git tags: ${currentVersion}"
                    
                    // Use npx semver to get the next version based on the current version
                    env.VERSION = sh(script: "npx semver -i patch ${currentVersion}", returnStdout: true).trim()
                    echo "Version set to: ${env.VERSION}"
                }
            }
        }

        stage('Login to Docker Hub') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: DOCKER_CREDENTIALS_ID, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh "echo '${DOCKER_PASSWORD}' | docker login -u '${DOCKER_USERNAME}' --password-stdin"
                    }
                }
            }
        }

        stage('Build and Push Multi-Arch Docker Image') {
            steps {
                script {
                    sh """
                        docker buildx build --platform linux/amd64,linux/arm64,linux/386,linux/ppc64le \\
                        -t ${IMAGE_NAME}:${TAG} . --push
                    """
                }
            }
        }
    }

    post {
        success {
            echo "Docker Image pushed successfully: ${IMAGE_NAME}:${TAG}"
        }
        failure {
            echo "Build failed!"
        }
    }
}
