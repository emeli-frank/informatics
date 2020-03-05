mkdir pdbqt

for %%f in (*.mol2) do (
    echo converting to pdbqt %%f
	
	obabel -imol2 %%f -opdbqt -Opdbqt\%%~nf.pdbqt
)

PAUSE

REM contact: dev.carlfrank@gmail.com