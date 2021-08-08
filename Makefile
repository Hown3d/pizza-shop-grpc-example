.PHONY: dev-up 

dev-up: 
	ctlptl apply -f ctlptl-cluster.yaml
	tilt up

dev-down:
	tilt down
	ctlptl delete -f ctlptl-cluster.yaml