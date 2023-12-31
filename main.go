package main

import (
	"fmt"
	"strategy-as1/config"
	"strategy-as1/factory"
	"strategy-as1/soldier_adapter"
	"strategy-as1/soldier_strategy"
)

func main() {
	//singleton
	instanceOfDB := config.GetInstanceOfDB()
	instanceOfDB.Info()
	configSingleton := config.GetInstanceOfDB()
	configSingleton.Info()
	//strategy/decorator
	basicSoldier := soldier_strategy.BasicSoldier{}
	bowSoldier := soldier_strategy.BowSoldier{Soldier: &basicSoldier}
	shieldSoldier := soldier_strategy.ShieldSoldier{Soldier: &basicSoldier}
	shieldBowSoldier := soldier_strategy.ShieldSoldier{
		Soldier: &soldier_strategy.BowSoldier{
			Soldier: &basicSoldier,
		},
	}
	fmt.Println()

	soldier1 := soldier_strategy.SoldierBehavior{SB: &basicSoldier}
	soldier2 := soldier_strategy.SoldierBehavior{SB: &bowSoldier}
	soldier3 := soldier_strategy.SoldierBehavior{SB: &shieldSoldier}
	soldier4 := soldier_strategy.SoldierBehavior{SB: &shieldBowSoldier}

	soldier1.DisplayStats()
	soldier2.DisplayStats()
	soldier3.DisplayStats()
	soldier4.DisplayStats()

	//adapter
	fmt.Println()
	adapter := &soldier_adapter.Adapter{}
	greekResident := &soldier_adapter.GreekResident{}
	adapter.TranslateTwoLanguages(greekResident)
	romanResident := &soldier_adapter.RomanResident{}
	romanTranslator := &soldier_adapter.RomanTranslator{
		RomanResident: romanResident,
	}
	adapter.TranslateTwoLanguages(romanTranslator)

	//factory
	fmt.Println()
	archer, _ := factory.GetSoldier("Archer")
	shieldbearer, _ := factory.GetSoldier("Shieldbearer")

	printDetails(archer)
	printDetails(shieldbearer)
}

//factory

func printDetails(i soldier_strategy.ISoldier) {
	fmt.Printf("Soldier: %s", i.GetName())
	fmt.Println()
	fmt.Printf("Attack: %d", i.GetAttack())
	fmt.Println()
	fmt.Printf("HP: %d", i.GetHP())
	fmt.Println()
}
