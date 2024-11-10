# Variables pour chaque service
API_SERVICE = api
GEOCODING_SERVICE = geocoding
FRONT_SERVICE = website

# Commande pour construire les images Docker de chaque projet
build: build-api build-geocoding build-front

build-api:
	@echo "Building API Docker image with existing command..."
	$(MAKE) -C $(API_SERVICE) dev

build-geocoding:
	@echo "Building Geocoding Docker image..."
	$(MAKE) -C $(GEOCODING_SERVICE) dev

build-front:
	@echo "Building Front Docker image..."
	$(MAKE) -C $(FRONT_SERVICE) dev

# Commande pour exécuter docker-compose
run:
	@echo "Running docker-compose..."
	docker-compose up
