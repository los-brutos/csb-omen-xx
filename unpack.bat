@echo off
REM 
REM dépaquetage des compendiums du module CSB Omen 12
REM le paquet fvtt doit être  installé https://github.com/foundryvtt/foundryvtt-cli
REM
REM /!\ nécessite que l'instance de foundry soit fermée avant de lancer la commande
REM
@echo off
cls
echo * ********************************** *
echo * start unpack compendiums           *
echo * ********************************** *
echo.
echo /// configure environment
call fvtt package workon csb-omen-xx --type=Module
echo.
echo /// unpack compendium "templates-acteurs"
call fvtt package unpack -c templates-acteurs
echo.
echo /// unpack compendium "templates-objets"
call fvtt package unpack -c templates-objets
echo.
echo /// unpack compendium "omen-systems"
call fvtt package unpack -c omen-systems
echo.
echo * ********************************** *
echo * all compendiums have been unpacked *
echo * ********************************** *
echo.
@echo on
