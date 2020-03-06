mkdir mol2

echo Now converting all smiles file (*.smi) in this directory to mol2
obabel *.smi -omol2 -m --gen3d

PAUSE

REM contact: dev.carlfrank@gmail.com