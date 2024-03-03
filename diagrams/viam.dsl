workspace {

    !identifiers hierarchical

    model {
        user = person "User"

        cliContainer = container "Order Servicec"
        appContainer = container "Customer Service"
        rdkContainer = container "Invoice Service"

        user -> appContainer 
        user -> cliContainer 
        cliContainer -> rdkContainer
        appContainer -> rdkContainer
    }

    views {
        styles {
            element "Person" {
                shape person
            }
        }
    }
        
}
