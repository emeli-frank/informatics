mkdir minimized

for %%f in (*.sdf) do (
    echo minimizing %%f
	
	obminimize -o sdf %%f > minimized\%%~nf.min.sdf
)

PAUSE

REM contact: dev.carlfrank@gmail.com