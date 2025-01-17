- Implement OpenAPI definitions
- Golang process on master node to call the rabbitMQ endpoints, fetch data and update CRDS
- Create Rancher extensions to handle sensors CRDs
- Create a manifest.yaml file for each supported sensors (svg, description, link to homepage)
- Sensor's CRDS should be updated ONLY by controller pod
- Rancher UI should display
  - sensor CRD
  - raspberry CRD
- Define a builder kit for devices
  - The device <-> rpc-server interface should use OpenAPI definition to build a skeleton and call the executable built in devices docker images.
  - The device's API should be defined in the settings file by dev
- Create Jobs to connect devices
- Define job python code using code-server in rancher extension 
- Inject python script in Job's ConfigMap