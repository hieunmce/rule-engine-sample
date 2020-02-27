package handler

import (
	"fmt"

	"github.com/hieunmce/rule-engine-sample/core"

	gruleAST "github.com/hyperjumptech/grule-rule-engine/ast"
	gruleBuilder "github.com/hyperjumptech/grule-rule-engine/builder"
	gruleEngine "github.com/hyperjumptech/grule-rule-engine/engine"
	grulePkg "github.com/hyperjumptech/grule-rule-engine/pkg"
)

func Handle() {
	patient := &core.Patient{
		Age:                2,
		Communication:      "Low",
		Socialization:      "High",
		DailyLivingSkill:   "High",
		DangerousBehaviors: "N",
	}

	dataCtx := gruleAST.NewDataContext()
	err := dataCtx.Add("Patient", patient)
	if err != nil {
		panic(err)
	}

	workingMemory := gruleAST.NewWorkingMemory()
	knowledgeBase := gruleAST.NewKnowledgeBase("tutorial", "1.0.0")
	ruleBuilder := gruleBuilder.NewRuleBuilder(knowledgeBase, workingMemory)

	bundle := grulePkg.NewGITResourceBundle("https://github.com/hieunmce/rule-engine-sample.git", "/**/*.grl")
	resources := bundle.MustLoad()
	for _, res := range resources {
		err := ruleBuilder.BuildRuleFromResource(res)
		if err != nil {
			panic(err)
		}
	}

	engine := gruleEngine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase, workingMemory)
	if err != nil {
		panic(err)
	}

	if patient.HaveRecommendation {
		// TODO
		fmt.Println("START TO INSERT RECOMMENDATION")
		fmt.Printf("%+v\n", patient)
	} else {
		fmt.Println("CAN'T MAKE RECOMMENDATION")
	}
}
