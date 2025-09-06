#!/bin/bash

# GoIt Core - Script de inicialização
set -e

echo "🚀 Iniciando GoIt Core com MongoDB..."

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Ir para o diretório raiz do projeto
cd "$(dirname "$0")/../.."

log_info "Parando containers existentes..."
sudo docker-compose down 2>/dev/null || true

log_info "Construindo imagem atualizada..."
sudo docker-compose build --no-cache goit-core

log_info "Iniciando serviços..."
sudo docker-compose up -d

log_info "Aguardando inicialização..."
sleep 5

# Verificar status
#if sudo docker ps --format "{{.Names}}" | grep -q "goit-core"; then
 #   log_success "GoIt Core rodando!"
#else
  #  log_error "Falha ao iniciar GoIt Core"
 #   exit 1
#fi

if sudo docker ps --format "{{.Names}}" | grep -q "mongo"; then
    log_success "MongoDB rodando!"
else
    log_error "Falha ao iniciar MongoDB"
    exit 1
fi

log_success "✅ Aplicação iniciada!"
log_info "📡 GoIt Core: localhost:50051"
log_info "🗄️  MongoDB: localhost:27017"
echo ""
log_info "Comandos úteis:"
log_info "  Ver logs: ./pkg/scripts/logs.sh"
log_info "  Parar: ./pkg/scripts/stop.sh"

