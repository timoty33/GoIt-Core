#!/bin/bash

# GoIt Core - Script para parar aplicação
set -e

# Cores
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Ir para o diretório raiz do projeto
cd "$(dirname "$0")/../.."

echo "⏹️  Parando GoIt Core..."

# Função de ajuda
show_help() {
    echo "Uso: ./pkg/scripts/stop.sh [opções]"
    echo ""
    echo "Opções:"
    echo "  -v, --volumes    Remover volumes também"
    echo "  -i, --images     Remover imagens também"
    echo "  -f, --force      Forçar parada (kill ao invés de stop)"
    echo "  -h, --help       Mostrar esta ajuda"
    echo ""
    echo "Exemplos:"
    echo "  ./pkg/scripts/stop.sh           # Parar containers normalmente"
    echo "  ./pkg/scripts/stop.sh -v        # Parar e remover volumes"
    echo "  ./pkg/scripts/stop.sh -f        # Forçar parada"
}

# Variáveis
REMOVE_VOLUMES=false
REMOVE_IMAGES=false
FORCE_STOP=false

# Parse dos argumentos
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--volumes)
            REMOVE_VOLUMES=true
            shift
            ;;
        -i|--images)
            REMOVE_IMAGES=true
            shift
            ;;
        -f|--force)
            FORCE_STOP=true
            shift
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        -*|--*)
            log_error "Opção desconhecida: $1"
            show_help
            exit 1
            ;;
        *)
            log_error "Argumento desconhecido: $1"
            show_help
            exit 1
            ;;
    esac
done

# Verificar se há containers rodando
RUNNING_CONTAINERS=$(sudo docker ps --format "{{.Names}}" | grep -E "goit-core|mongo" || true)

if [[ -z "$RUNNING_CONTAINERS" ]]; then
    log_info "Nenhum container do GoIt Core está rodando."
    exit 0
fi

log_info "Containers em execução:"
echo "$RUNNING_CONTAINERS" | sed 's/^/  /'

# Parar containers
if [[ "$FORCE_STOP" == true ]]; then
    log_warning "Forçando parada dos containers..."
    sudo docker-compose kill
else
    log_info "Parando containers graciosamente..."
    sudo docker-compose down
fi

# Remover volumes se solicitado
if [[ "$REMOVE_VOLUMES" == true ]]; then
    log_warning "Removendo volumes..."
    sudo docker-compose down --volumes
    log_info "⚠️  Dados do MongoDB foram removidos!"
fi

# Remover imagens se solicitado
if [[ "$REMOVE_IMAGES" == true ]]; then
    log_warning "Removendo imagens..."
    
    # Listar imagens relacionadas ao projeto
    PROJECT_IMAGES=$(sudo docker images --format "{{.Repository}}:{{.Tag}}" | grep -E "goit.*core|mongo" || true)
    
    if [[ -n "$PROJECT_IMAGES" ]]; then
        log_info "Removendo imagens:"
        echo "$PROJECT_IMAGES" | sed 's/^/  /'
        echo "$PROJECT_IMAGES" | xargs sudo docker rmi -f 2>/dev/null || true
    fi
fi

# Verificar se parou com sucesso
sleep 2
STILL_RUNNING=$(sudo docker ps --format "{{.Names}}" | grep -E "goit-core|mongo" || true)

if [[ -z "$STILL_RUNNING" ]]; then
    log_success "✅ Todos os containers foram parados!"
else
    log_error "Alguns containers ainda estão rodando:"
    echo "$STILL_RUNNING" | sed 's/^/  /'
    exit 1
fi

echo ""
log_info "Para iniciar novamente: ./pkg/scripts/start.sh"

# Mostrar resumo do que foi feito
echo ""
log_info "📊 Resumo:"
if [[ "$FORCE_STOP" == true ]]; then
    echo "  🔴 Containers foram forçadamente parados"
else
    echo "  ⏹️  Containers foram parados graciosamente"
fi

if [[ "$REMOVE_VOLUMES" == true ]]; then
    echo "  🗑️  Volumes foram removidos"
fi

if [[ "$REMOVE_IMAGES" == true ]]; then
    echo "  🖼️  Imagens foram removidas"
fi
