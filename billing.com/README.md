## Getting Started
First, run on the development server:

```bash
npm i && npm run dev
```

## Run on Docker
Don't forget to setup the .env file before build the docker image

```bash
npm i && npm run build 
docker-compose up
```

## Environment Variable
NEXT_PUBLIC_API_HOST, default = http://localhost:38000. Specifies which backend server the system uses.
