package display

import (
	"FilesDIR/globals"
	"fmt"
	"path/filepath"
	"time"
)

func DrawStart() string {
	return fmt.Sprintf(`<fg=48,207,37>
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝</>
		<cyan>Version:</> <fg=48,207,37>%s</>              <cyan>Auteur:</> <fg=48,207,37>%s</>


`, globals.Version, globals.Author)
}

func DrawInitSearch() string {
	return fmt.Sprint(`<fg=214,99,144>Initialisation du programme...</>`)
}

func DrawRunSearch() string {
	return fmt.Sprint(`<cyan>
+============================================================+
|                    DEBUT DES RECHERCHES                    |
+============================================================+
</>`)
}

func DrawEndSearch() string {
	return fmt.Sprint(`<cyan>
+============================================================+
|                     FIN DES RECHERCHES                     |                      
+============================================================+
</>`)
}

func DrawWriteExcel() string {
	return fmt.Sprint(`<fg=214,99,144>Export Excel...   `)
}

func DrawSaveExcel() string {
	return fmt.Sprint(`<fg=214,99,144>Fichier Excel sauvegardé avec succes.</>`)
}

func DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles, PoolSize int, timerSearch time.Duration, timerTotal time.Duration) string {
	return fmt.Sprintf(`<cyan>
+============================================================+
|                    BILAN DES RECHERCHES                    |                     
+============================================================+
</>
<fg=214,99,144>#### - INFOS GENERALES :</>
<cyan>Dossiers principal:</> <green>%s</>
<cyan>Requête utilisée:</> <green>%s</>
<cyan>Nombre de Threads:</> <green>%v</>
<cyan>Nombre de Goroutines:</> <green>%v</>

<bg=255,35,156>#### - RESULTATS :</>
<cyan>Fichiers trouvés:</> <green>%v</>
<cyan>Temps d'exécution de la recherche:</> <green>%v</>
<cyan>Temps d'exécution total:</> <green>%v</>

<fg=214,99,144>#### - EXPORTS :</>
<cyan>Logs:</> <green>%s</>
<cyan>Dumps:</> <green>%s</>
<cyan>Export Excel:</> <green>%s</>

<cyan>+=========  Auteur:</> <yellow>%s</>       <cyan>Version:</> <yellow>%s</>  <cyan>=========+</>
`,
		SrcPath,
		ReqFinal,
		PoolSize,
		NbGoroutine,

		NbFiles,
		timerSearch,
		timerTotal,

		filepath.Join(globals.TempPathGen, "logs"),
		filepath.Join(globals.TempPathGen, "dumps"),
		DstPath,

		globals.Author, globals.Version)
}

func DrawInitCompiler() string {
	return fmt.Sprint(`Initialisation de la compilation...`)
}

func DrawRunCompiler() string {
	return fmt.Sprint(`
+============================================================+
|                   DEBUT DES COMPILATIONS                   |
+============================================================+
`)
}

func DrawEndCompiler() string {
	return fmt.Sprint(`
+============================================================+
|                    FIN DES COMPILATIONS                    |
+============================================================+
`)
}
