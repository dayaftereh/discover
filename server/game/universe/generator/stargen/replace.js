const fs = require("fs")

const content = fs.readFileSync('./gas-radius.go')
var data = content.toString()

data = data.replace(/planet_radius_helper/g, 'planetRadiusHelper')
data = data.replace(/total_earth_masses/g, 'totalEarthMasses')
data = data.replace(/core_earth_masses/g, 'coreEarthMasses')
data = data.replace(/jupiter_radii/g, 'jupiterRadii')
data = data.replace(/jupiterRadii1 =/g, 'jupiterRadii1 :=')
data = data.replace(/jupiterRadii2 =/g, 'jupiterRadii2 :=')

var replaceCoreMess = n =>{
    const regex = new RegExp(`mass_radii\\[${n}\\]`,'g')
    data = data.replace(regex, `massRadii${n}`)
}

//replaceCoreMess(10)
//replaceCoreMess(25)
//replaceCoreMess(50)
replaceCoreMess(100)

data = data.replace(/core_mass_radii\[100\]/g, 'coreMassRadii100')
data = data.replace(/core_mass_radii\[0\]/g, 'coreMassRadii0')
data = data.replace(/core_mass_radii\[10\]/g, 'coreMassRadii10')
data = data.replace(/core_mass_radii\[25\]/g, 'coreMassRadii25')
data = data.replace(/core_mass_radii\[50\]/g, 'coreMassRadii50')


data = data.replace(/coreMassRadii100 =/g, 'coreMassRadii100 :=')
data = data.replace(/coreMassRadii0 =/g, 'coreMassRadii0 :=')
data = data.replace(/coreMassRadii10 =/g, 'coreMassRadii10 :=')
data = data.replace(/coreMassRadii25 =/g, 'coreMassRadii25 :=')
data = data.replace(/coreMassRadii50 =/g, 'coreMassRadii50 :=')


var massRadii = `
gas_300Myr_78K_100core_mass[129] = 0.587;
gas_300Myr_78K_100core_mass[215] = 0.81;
gas_300Myr_78K_100core_mass[318] = 0.92;
gas_300Myr_78K_100core_mass[464] = 0.999;
gas_300Myr_78K_100core_mass[774] = 1.072;
gas_300Myr_78K_100core_mass[1292] = 1.107;
gas_300Myr_78K_100core_mass[2154] = 1.119;
gas_300Myr_78K_100core_mass[3594] = 1.107;
`
var list = massRadii.split(/\n|\r\n/g)

list.forEach(line => {
    line = line.trim()
    if (!line || line.length < 1) {
        return
    }
    const values = line.split("=", 2)
    if (!values || values.length < 2) {
        return
    }
    var x = values[1].trim()

    if (x.endsWith(";")) {
        x = x.replace(";", "")
    }

    var name = values[0]
    const result = name.match(/.*\[(\d+)\].*/)


    if (!result || result.length < 2) {
        return
    }
    var num = result[1]
    //console.log(`mass_radii[${num}]`)
    const regex = new RegExp(`mass_radii\\[${num}\\]`, 'g')
    data = data.replace(regex, x)

})

fs.writeFileSync("./gas-radius.go", data)
