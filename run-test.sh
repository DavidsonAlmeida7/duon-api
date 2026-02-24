#!/bin/bash

DIR=${1:-./...}

GREEN="\e[32m"
RED="\e[31m"
RED_LIGHT="\e[91m"
YELLOW="\e[33m"
BLUE="\e[34m"
MAGENTA="\e[35m"
CYAN="\e[36m"
RESET="\e[0m"

FAILED_TESTS_FILE=$(mktemp)
ERROR_DETAILS_FILE=$(mktemp)
FULL_ERROR_FILE=$(mktemp)

clean_test_name() {
    echo "$1" | sed -E 's/^Test//' | sed 's/([0-9.]*s)//g' | xargs
}

clear
echo -e "\n===== Executando Testes =====${RESET}\n"

current_test=""
error_section=false

docker exec -i emcash_simulador go test -v $DIR -cover | while IFS= read -r line; do
    if [[ "$line" == *"=== RUN"* ]]; then
        current_test=$(echo "$line" | cut -d' ' -f4-)
        error_section=false

        continue
    fi

    if [[ "$line" == *"--- PASS:"* ]]; then
        TEST_NAME=$(clean_test_name "$line")

        echo -e "         ${GREEN} 🗸 ${RESET}${line#--- PASS: }  ${RESET}"

        continue
    fi

    if [[ "$line" == *"--- FAIL:"* ]]; then
        TEST_NAME=$(clean_test_name "$line")

        echo -e "         ${RED} 🗙 ${RED_LIGHT}${line#--- FAIL: } ${RESET}"
        echo "$current_test" >> "$FAILED_TESTS_FILE"

        error_section=true

        continue
    fi


    if [[ "$line" == *"ok"* ]]; then
        PACKAGE_INFO="${line#ok }"

        echo -e "Pacote:      $PACKAGE_INFO"

        continue
    fi

done

if [[ -s "$FAILED_TESTS_FILE" ]]; then
    echo -e "\n===== TESTES QUEBRADOS =====${RESET}\n"
    
    count=1
    
    while IFS= read -r failed_test; do
        error_trace=$(grep "Error Trace:" "$FULL_ERROR_FILE" | grep "$failed_test" | head -n 1)
        error_message=$(grep "Error:" "$FULL_ERROR_FILE" | grep "$failed_test" | head -n 1)
        additional_messages=$(grep "Messages:" "$FULL_ERROR_FILE" | grep "$failed_test")
        
        echo -e "${count}) ${RED}🗙 ${RESET}${failed_test}${RESET}"
        ((count++))
    done < "$FAILED_TESTS_FILE"
else
    echo -e "\n${GREEN}Todos os testes passaram! 🚀${RESET}"
fi

rm -f "$FAILED_TESTS_FILE" "$ERROR_DETAILS_FILE" "$FULL_ERROR_FILE"
