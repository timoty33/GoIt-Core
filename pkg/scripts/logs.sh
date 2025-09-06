#!/bin/bash

# GoIt Core - Script para visualizar logs

# Cores
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }

# Ir para o diretório raiz do projeto
cd "$(dirname "$0")/../.."

# Função de ajuda
show_help() {
    echo "Uso: ./pkg/scripts/logs.sh [opções] [serviço]"
    echo ""
    echo "Serviços disponíveis:"
    echo "  goit-core  - Logs da aplicação GoIt Core"
    echo "  mongo      - Logs do MongoDB"
    echo "  all        - Logs de todos os serviços (padrão)"
    echo ""
    echo "Opções:"
    echo "  -f, --follow   Seguir logs em tempo real"
    echo "  -n, --lines N  Mostrar últimas N linhas (padrão: 50)"
    echo "  -h, --help     Mostrar esta ajuda"
    echo ""
    echo "Exemplos:"
    echo "  ./pkg/scripts/logs.sh                    # Ver todos os logs"
    echo "  ./pkg/scripts/logs.sh goit-core          # Apenas logs da app"
    echo "  ./pkg/scripts/logs.sh -f mongo           # Seguir logs do mongo"
    echo "  ./pkg/scripts/logs.sh -n 100 goit-core   # Últimas 100 linhas"
}

""" Verificar se containers estão rodando
check_containers() {
    if ! sudo docker ps --format "{{.Names}}" | grep -q "goit-core\|mongo"; then
        log_info "⚠️  Nenhum container está rodando."
        log_info "Para iniciar: ./pkg/scripts/start.sh"
        exit 1
    fi
}
"""

# Variáveis padrão
FOLLOW=false
LINES=50
SERVICE=""

# Parse dos argumentos
while [[ $# -gt 0 ]]; do
    case $1 in
        -f|--follow)
            FOLLOW=true
            shift
            ;;
        -n|--lines)
            LINES="$2"
            shift 2
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        goit-core|mongo|all)
            SERVICE="$1"
            shift
            ;;
        -*|--*)
            echo "Opção desconhecida: $1"
            show_help
            exit 1
            ;;
        *)
            SERVICE="$1"
            shift
            ;;
    esac
done

# Serviço padrão
if [[ -z "$SERVICE" ]]; then
    SERVICE="all"
fi

check_containers

echo "📋 Visualizando logs..."

# Construir comando docker-compose logs
DOCKER_CMD="sudo docker-compose logs"

if [[ "$FOLLOW" == true ]]; then
    DOCKER_CMD="$DOCKER_CMD --follow"
fi

DOCKER_CMD="$DOCKER_CMD --tail $LINES"

# Adicionar timestamps e cores
DOCKER_CMD="$DOCKER_CMD --timestamps"

case $SERVICE in
    "goit-core")
        log_info "🚀 Logs do GoIt Core:"
        $DOCKER_CMD goit-core
        ;;
    "mongo")
        log_info "🗄️ Logs do MongoDB:"
        $DOCKER_CMD mongo
        ;;
    "all")
        log_info "📊 Logs de todos os serviços:"
        $DOCKER_CMD
        ;;
    *)
        echo "Serviço desconhecido: $SERVICE"
        show_help
        exit 1
        ;;
esac
