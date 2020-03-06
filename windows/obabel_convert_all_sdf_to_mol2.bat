mkdir mol2

for %%f in (*.sdf) do (
    echo converting to mol2 %%f
	
	obabel -i sdf %%f -o mol2 -O mol2\%%~nf.mol2
)

PAUSE

REM contact: dev.carlfrank@gmail.com