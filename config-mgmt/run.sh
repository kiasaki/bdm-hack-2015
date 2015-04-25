export ANSIBLE_LIBRARY="$ANSIBLE_LIBRARY:library"
ansible-playbook playbook.yml -i inventory.ini --private-key=~/.ssh/bdmhack --vault-password-file=~/bdmhack-vp.txt $ARGS
