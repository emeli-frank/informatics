mkdir minimized

for %%f in (*.mol2) do (
    echo minimizing %%f
	
	obminimize -o mol2 %%f > minimized\%%~nf.min.mol2
)

PAUSE

REM contact: dev.carlfrank@gmail.com