# usando docker

```docker run --rm -it -p 8089:80 -v /home/rancher/admin:/usr/share/nginx -w /usr/share/nginx dsaouda/nginx bash -c "cd html && git pull origin master && nginx -g 'daemon off;'"```

# usando Jenkins

```bash

#!/bin/bash

echo "pegando id do container"

# pegando id container
id=`curl -s -u "E976925C82406A5E9796:iNVcLfVWjFaFsU6pkW4GWhXoSo7yjCurs6S4QTBT" -X GET http://192.168.200.10:8080/v2-beta/projects/1a5/container?name_eq=adminui | cut -d ':' -f 20 | cut -d '"' -f 2`


echo "removendo container $id"

# remover o container antigo
curl -u "E976925C82406A5E9796:iNVcLfVWjFaFsU6pkW4GWhXoSo7yjCurs6S4QTBT" \
-X POST \
-H 'Content-Type: application/json' \
-d '{
	"remove": true,
	"timeout": 0
}' http://192.168.200.10:8080/v2-beta/projects/1a5/containers/$id?action=stop


echo "aguardando container ser removido"
sleep 5

echo "criando container"

# criando container
curl -s -u "E976925C82406A5E9796:iNVcLfVWjFaFsU6pkW4GWhXoSo7yjCurs6S4QTBT" \
-X POST \
-H 'Content-Type: application/json' \
-d '{"instanceTriggeredStop":"stop","startOnCreate":true,"publishAllPorts":false,"privileged":false,"stdinOpen":true,"tty":true,"readOnly":false,"runInit":false,"networkMode":"bridge","type":"container","requestedHostId":"1h6","secrets":[],"dataVolumes":["/home/rancher/admin:/usr/share/nginx"],"dataVolumesFrom":[],"dns":[],"dnsSearch":[],"capAdd":[],"capDrop":[],"devices":[],"logConfig":{"driver":"","config":{}},"dataVolumesFromLaunchConfigs":[],"imageUuid":"docker:dsaouda/nginx","ports":["8089:80/tcp"],"instanceLinks":{},"labels":{},"name":"adminui","workingDir":"/usr/share/nginx","networkContainerId":null,"hostname":"adminui","command":["bash","-c","cd html && git pull origin master && nginx -g \"daemon off;\""],"count":null,"createIndex":null,"created":null,"deploymentUnitUuid":null,"description":null,"externalId":null,"firstRunning":null,"healthState":null,"kind":null,"memoryReservation":null,"milliCpuReservation":null,"removed":null,"startCount":null,"uuid":null,"volumeDriver":null,"user":null,"domainName":null,"memorySwap":null,"memory":null,"cpuSet":null,"cpuShares":null,"pidMode":null,"blkioWeight":null,"cgroupParent":null,"usernsMode":null,"pidsLimit":null,"diskQuota":null,"cpuCount":null,"cpuPercent":null,"ioMaximumIOps":null,"ioMaximumBandwidth":null,"cpuPeriod":null,"cpuQuota":null,"cpuSetMems":null,"isolation":null,"kernelMemory":null,"memorySwappiness":null,"shmSize":null,"uts":null,"ipcMode":null,"stopSignal":null,"oomScoreAdj":null,"ip":null,"ip6":null,"healthInterval":null,"healthTimeout":null,"healthRetries":null}' \
http://192.168.200.10:8080/v2-beta/projects/1a5/container

```
