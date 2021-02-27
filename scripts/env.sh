if [ ! -f .env ]; then
    echo ".env not found"
    exit 1
fi

export $(cat .env | grep -v '^#')
exec $@