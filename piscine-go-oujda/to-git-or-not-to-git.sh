
curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .appearance.gender, .powerstats.power'