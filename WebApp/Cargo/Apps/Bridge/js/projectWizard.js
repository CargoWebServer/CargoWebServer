/**
 * The project wizard.
 */
var ProjectWizard = function(){

    Wizard.call(this, "project_wizard", "New Project", 400, 400)

    // Now I will create the pages...
    //var templatePageContent = new Element(null, {"tag":"div", "style":"width: 100%; height:100%;"})
    //var templatesPage = new WizardPage("wizard_template_page", "Project Type", "Chose a template for your new project.", templatePageContent, 600, 400)
    //this.appendPage(templatesPage)

    var page1 = new WizardPage("test_page_1", "Test 1", "This is a test 1!", new Element(null, {"tag":"div", "style":"width: 100%; height:100%; ", "innerHtml": "page 1"}), 600, 400)
    var page2 = new WizardPage("test_page_2", "Test 2", "This is a test 2!", new Element(null, {"tag":"div", "style":"width: 100%; height:100%; ", "innerHtml": "page 2"}), 600, 400)
    var page3 = new WizardPage("test_page_3", "Test 3", "This is a test 3!", new Element(null, {"tag":"div", "style":"width: 100%; height:100%; ", "innerHtml": "page 3"}), 600, 400)
    var page4 = new WizardPage("test_page_4", "Test 4", "This is a test 4!", new Element(null, {"tag":"div", "style":"width: 100%; height:100%; ", "innerHtml": "page 4"}), 600, 400)

    this.appendPage(page1)
    this.appendPage(page2)
    this.appendPage(page3)
    this.appendPage(page4)

    return this
}

ProjectWizard.prototype = new Wizard(null);
ProjectWizard.prototype.constructor = ProjectWizard;
