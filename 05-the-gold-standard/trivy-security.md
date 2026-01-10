# Trivy Vulnerability Scanning Configuration
# Add this to your Level 5 workflow

# scan-config.yaml
# Use this when running 'trivy config .' or 'trivy image gold-app'

severity: 'HIGH,CRITICAL'
exit-code: 1 # Fail the build if High/Critical vulnerabilities are found
ignore-unfixed: true

# Example CI Command:
# trivy image --severity HIGH,CRITICAL --exit-code 1 --ignore-unfixed gold-app:v1
