Installe Git : https://git-scm.com
Ouvre le lien (CTRL+CLICK GAUCHE sur le lien)
CLique sur install for Windows
Accepte tout et installe

Installe VScode : https://code.visualstudio.com
Ouvre le lien (CTRL+CLICK GAUCHE sur le lien)
CLique sur Download for Windows
Accepte tout et installe

Installe Golang : https://go.dev
Ouvre le lien (CTRL+CLICK GAUCHE sur le lien)
CLique sur Download for Windows
Accepte tout et installe

copie le lien du GitHub

fait Win+R
écrit "cmd"
fait entrez
saisi "git config --global user.name [Ton username]"
fait entrez
saisi "git config --global user.email [Ton addresse mail]"
fait entrez
Ferme l'onglet "Invite de Commande"

fait win+R
écrit "cmd"
fait entrez
saisi "cd Desktop"
fait entrez
saisi "mkdir Projet-ultra-top-secret-vraiment-maximum-top-confidentiel"
fait entrez
saisi "cd Projet-ultra-top-secret-vraiment-maximum-top-confidentiel"
fait entrez
saisi "mkdir Le-Projet-Ultra-confidentiel-va-se-retrouver-dans-ce-dossier"
fait entrez
saisi "cd Le-Projet-Ultra-confidentiel-va-se-retrouver-dans-ce-dossier"
fait entrez
saisi"mkdir Cest-ici-que-le-fichier-ce-trouve-car-cest-important"
fait entrez
saisi "cd Cest-ici-que-le-fichier-ce-trouve-car-cest-important"
fait entrez
saisi "git clone [lien du github]"
fait entrez

Fait Win+R
écrit "cmd"
fait entrez
saisi "cd Desktop/Projet-ultra-top-secret-vraiment-maximum-top-confidentiel/Le-Projet-Ultra-confidentiel-va-se-retrouver-dans-ce-dossier/Cest-ici-que-le-fichier-ce-trouve-car-cest-important"
fait entrez
saisi "code ."
fait entrez

Faire CTRL + K + T
Sélectionner Light (Visual Studio) Visual Studio Light
Aller dans le dossier a gauche nommé fichiertxt
cliquer sur Sample.txt
écrit ce que tu souhaites en connaissant les consignes si dessous(Ligne:80)

Ou Alors crée ton .txt avec ce que tu veux écrire dedans avec les mêmes consignes
va dans tes dossier
va dans le chemin d'accès Desktop (bureau)
puis dans Projet-ultra-top-secret-vraiment-maximum-top-confidentiel
ensuite dans Le-Projet-Ultra-confidentiel-va-se-retrouver-dans-ce-dossier
après dans Cest-ici-que-le-fichier-ce-trouve-car-cest-important
enfin dans fichiertxt
Et ici déposer votre fichier.txt

retourner sur votre VScode
cliquer sur Terminal en haut a gauche
cliquer sur New Terminal
saisi "go run main.go 
fait entrez
saisi [Le nom du fichier.txt]" (Sample.txt si vous avez choisi d'écrire dans le Sample.txt ou alors le nom de votre fichier que vous avez déposé précedement)
fait entrez

Votre texte avec les modifications appliqué apparaitrera dans le fichier texte Result.txt

Consigne :
dans votre texte si vous mettez (hex) : ce qu'il y aura écrit devant devra etre de l'hexadecimal et le fait d'écrire le (hex) va transformer le nombre devant et le transformer en décimal , pareil pour le binaire mais avec (bin)
Si vous utiliser (up) ou (low) ou (cap) cela va transformer le mot avant en, pour (up) mettre le mot tout en majuscule, pour le (low) tout en minuscule , et pour le (cap) ca va mettre la premiere lettre en majuscule.
également le programme va formater les caracteres comme les virgules pour les coller au mot précedent (exemple: "test , test" --> "test, test")
également le programme va formater les caracteres comme les apostrophes pour les coller au mot encadrer (exemple: "test ' test test ' test" --> "test 'test test' test")