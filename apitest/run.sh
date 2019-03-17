#!/bin/bash

# Check if VAULT_ADDR is set
if [ ! -z ${VAULT_ADDR} ]; then
    if [ -z ${VAULT_TOKEN} ]; then
        KUBE_TOKEN=$(</var/run/secrets/kubernetes.io/serviceaccount/token)

        RESULT=$(curl -sb  --request POST --data "{\"role\": \"${ROLE}\", \"jwt\": \"$KUBE_TOKEN\" }" "$VAULT_ADDR/v1/auth/kubernetes/login")
        TOKEN=$(echo $RESULT | jq -r .auth.client_token)
    else
        TOKEN=${VAULT_TOKEN}
    fi
    kv=$(curl --silent --header "X-Vault-Token: $TOKEN" "${VAULT_ADDR}/v1/secret/microservices/configuration/${VAULT_PATH}" | jq -r '.data')
    # Get json key from key/value
    keys=$(jq 'keys[]'<<<$kv)



    for item in $keys; do
        # Get value from selected key
            values=$(jq ".$item"<<<$kv)
        # Remove double quote from all variable
        key=$(echo $item | tr -d '"')
        value=$(echo $values | tr -d '"' )
            # Sed-ing to config
        sed -i "s|{{ $key }}|$value|g" /opt/microservice_distributor_deposit/configurations/App.yaml
    done


fi
