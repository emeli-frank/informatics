mkdir result

for %%f in (ligands\*.pdbqt) do (
	mkdir result\%%~nf
	vina --config conf.txt --ligand ligands\%%~nf.pdbqt --out result\%%~nf\out.pdbqt --log result\%%~nf\log.txt
)

PAUSE

REM contact: carlfrankben@gmail.com