#!/bin/bash

# build environment replacements
sed -i "s/NEWRELICLICENSEKEY/$NEWRELIC_LICENSE_KEY/g" ./manifest.yml
sed -i "s/SPACE/$SPACE/g" ./manifest.yml
sed -i "s/APPLICATION/$APPLICATION/g" ./manifest.yml
sed -i "s/MAJOR/$MAJOR/g" ./manifest.yml
sed -i "s/MINOR/$MINOR/g" ./manifest.yml
sed -i "s/TRAVIS_BUILD_NUMBER/$TRAVIS_BUILD_NUMBER/g" ./manifest.yml

# cf deployment
cf login -a "${1}" -o "${2}" -s "${3}" -u "${4}" -p "${5}"
cf push

# map new route if app is running
if [ $(curl http://$SPACE-$APPLICATION-blue.$DOMAIN/ping -s) == "OK" ]
then
    cf map-route $SPACE-$APPLICATION-$MAJOR.$MINOR.$TRAVIS_BUILD_NUMBER $DOMAIN --hostname $SPACE-$APPLICATION
    if [ "${SPACE}" == "prod" ]
    then
        # create domains and map route
        for tawerin_domain in "${TAWERIN_DOMAINS[@]}"
        do
            echo "creating and mapping ${tawerin_domain} to ${SPACE}-${APPLICATION}-${MAJOR}.${MINOR}.${TRAVIS_BUILD_NUMBER}"
            cf create-domain $BM_ORG $tawerin_domain
            cf map-route $SPACE-$APPLICATION-$MAJOR.$MINOR.$TRAVIS_BUILD_NUMBER $tawerin_domain
        done
    fi
fi

# delete old cf instances once app if app is running
if [ $(curl http://$SPACE-$APPLICATION.$DOMAIN/ping -s) == "OK" ]
then
    echo "${SPACE}-${APPLICATION} is RUNNING"
    for application in $(cf a | grep $APPLICATION | grep -v $APPLICATION-$MAJOR.$MINOR.$TRAVIS_BUILD_NUMBER | awk '{print $1}')
    do
        echo "deleting ${application}"
        cf delete -f $application
    done
    cf unmap-route $SPACE-$APPLICATION-$MAJOR.$MINOR.$TRAVIS_BUILD_NUMBER $DOMAIN --hostname $SPACE-$APPLICATION-blue
else
    echo "${SPACE}-${APPLICATION} is not RUNNING"
fi
