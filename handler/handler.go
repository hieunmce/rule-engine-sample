package handler

import (
	"errors"
	"fmt"

	"github.com/hieunmce/rule-engine-sample/core"

	gruleAST "github.com/hyperjumptech/grule-rule-engine/ast"
	gruleBuilder "github.com/hyperjumptech/grule-rule-engine/builder"
	gruleEngine "github.com/hyperjumptech/grule-rule-engine/engine"
	grulePkg "github.com/hyperjumptech/grule-rule-engine/pkg"
)

func HandleEnqueue(patient *core.Patient) error {
	dataCtx := gruleAST.NewDataContext()
	err := dataCtx.Add("Patient", patient)
	if err != nil {
		return err
	}

	workingMemory := gruleAST.NewWorkingMemory()
	knowledgeBase := gruleAST.NewKnowledgeBase("tutorial", "1.0.0")
	ruleBuilder := gruleBuilder.NewRuleBuilder(knowledgeBase, workingMemory)

	bundle := grulePkg.NewGITResourceBundle("https://github.com/hieunmce/rule-engine-sample.git", "/**/*.grl")
	resources := bundle.MustLoad()
	for _, res := range resources {
		err := ruleBuilder.BuildRuleFromResource(res)
		if err != nil {
			return err
		}
	}

	engine := gruleEngine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase, workingMemory)
	if err != nil {
		return err
	}

	if patient.HaveRecommendation {
		fmt.Println("START TO INSERT RECOMMENDATION")
		fmt.Printf("%+v\n", patient)
		return nil
	}
	return errors.New("CAN'T MAKE RECOMMENDATION")
}
