#!/bin/bash -e

if [[ -z "${KUBECONFIG}" ]]; then
  echo "Please set the [KUBECONFIG] to a valid file path"
  exit 1
fi


LOGDIR="cse-k8s-tkgm-logs-"$(date '+%F-%H-%M-%S')
mkdir -p ${LOGDIR}

function logger() {
  echo "-------- START LOG $2" >> "${LOGDIR}/$1"
  /bin/bash -c "$2" &>> "${LOGDIR}/$1"
  exitCode="$?"

  # add a newline in case the command does not add it
  last=$(tail -c 1 "${LOGDIR}/$1")
  if [ -n "$last" ]; then
    echo >> "${LOGDIR}/$1"
  fi

  echo "-------- END   LOG $2 CODE $exitCode" >> "${LOGDIR}/$1"
}

echo "Collecting Kubernetes details..."
logger kubernetes.txt 'kubectl version'
logger kubernetes.txt 'kubectl get nodes -owide'
logger kubernetes.txt 'kubectl get pods -A -owide'
logger kubernetes.txt 'kubectl get events -A'
logger kubernetes.txt 'kubectl describe nodes'
logger kubernetes.txt 'kubectl describe pods -A'
logger kubernetes.txt 'kubectl describe deployments -A'
logger kubernetes.txt 'kubectl describe replicasets -A'
logger kubernetes.txt 'kubectl describe services -A'
logger kubernetes.txt 'kubectl describe configmaps -A'
logger kubernetes.txt 'kubectl get secrets -A'
logger kubernetes.txt 'kubectl describe pv -A'


echo "Collecting Kubernetes logs..."
kubectl cluster-info dump -A --output-directory="${LOGDIR}"/k8s-cluster-info &>/dev/null

echo "Compressing logs and generating tarball at ${LOGDIR}.tar.gz..."
tar -czvf "${LOGDIR}.tar.gz" "${LOGDIR}" &> /dev/null


exit 0
