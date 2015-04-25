import sys
import json

STATE_FILE = 'terraform.tfstate'
INVENTORY_FILE = 'inventory.ini'
INVENTORY_TEMPLATE = '{0}.tmpl'.format(INVENTORY_FILE)


def read_and_parse_state(state_file):
    state_file_contents = None
    with open(state_file, 'r') as f:
        state_file_contents = f.read()
    return json.loads(state_file_contents)


def ips_from_state(state):
    return state.get('modules')[0].get('outputs')


def format_inventory_template(inventory_template, values):
    with open(inventory_template, 'r') as template:
        return template.read().format(**values)


def main():
    # Honor cli args
    state_file = STATE_FILE
    if len(sys.argv) >= 2:
        state_file = sys.argv[1]

    inventory_template = INVENTORY_TEMPLATE
    if len(sys.argv) >= 3:
        inventory_template = sys.argv[2]

    print('Reding state from: {}'.format(state_file))
    state = read_and_parse_state(state_file)

    ips = ips_from_state(state)
    print('Found outputs: {}'.format(ips))

    formatted_inventory_contents = format_inventory_template(
        inventory_template, ips)

    with open(INVENTORY_FILE, 'w') as target:
        target.truncate()
        target.write(formatted_inventory_contents)
        target.flush()

    print('All done!')


main()
