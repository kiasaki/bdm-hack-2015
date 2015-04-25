default:
	@echo "plan, apply, plan-destroy, update-inventory"

plan:
	terraform plan -out terraform.tfplan

plan-destroy:
	terraform plan -out terraform.tfplan -destroy

apply:
	terraform apply terraform.tfplan

update-inventory:
	(cd  ../config-mgmt && python3 inventory.py ../infrastructure/terraform.tfstate)
