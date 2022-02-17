package main

import (
	"fmt"
	"strings"
)

func FormatLDAP( whitelist []string) {

    for _, value := range whitelist {
         
        // Split the string into slices separated by comma
        splitValue := strings.Split(value, ",")
        // Send each attribute through the necessary checks
        for _, singleAttribute := range splitValue {
            err := CheckPrefix(singleAttribute)
            if err != nil {
                fmt.Printf("Attribute formatting error: %s\n", singleAttribute)
            }
            if strings.HasPrefix(singleAttribute, "CN=") {
                err := CheckCN(singleAttribute)
                if err != nil {
                    fmt.Printf("Common name contains invalid character: %s\n", singleAttribute)
                }
            }
            if strings.HasPrefix(singleAttribute, "OU=") {
                err := CheckOU(singleAttribute)
                if err != nil {
                    fmt.Printf("OU name error: %s\n", singleAttribute)
                }
            }
            if strings.HasPrefix(singleAttribute, "DC=") {
                err := CheckDC(singleAttribute)
                if err != nil {
                    fmt.Printf("Domain component error: %s\n", singleAttribute)
                }
            }
        }
   
    }
}

// CheckPrefixis a function that checks if the query contains only CN, OU, and DC, err if not
func CheckPrefix(whitelist_query string) error {

    attribute := strings.HasPrefix(whitelist_query, "CN=") || strings.HasPrefix(whitelist_query, "DC=") || strings.HasPrefix(whitelist_query, "OU=")     
    switch attribute {
    case true:
        return nil
    default:
        return fmt.Errorf("Incorrect attribute naming convention")
    }
}
func CheckCN(common_name string) error {
    cn_name := strings.ContainsAny(common_name, "+)(*&^%$#@!")
    switch cn_name {
    case true:
        return fmt.Errorf("Incorrect naming convention for common name")
    default:
        return nil
    }
}

// CheckOU is a function that checks if the query contains the right organization unit options, err if not
func CheckOU(organizational_unit string) error {
    
    ou_name := strings.HasSuffix(organizational_unit,"=DOCET") || strings.HasSuffix(organizational_unit,"=Admin Groups") || strings.HasSuffix(organizational_unit,"=Admin") || 
    strings.HasSuffix(organizational_unit,"=Infrastructure") || strings.HasSuffix(organizational_unit,"=Security Groups") || strings.HasSuffix(organizational_unit,"=Worldpay")

    switch ou_name {
    case true:
        return nil
    default:
        return fmt.Errorf("Organizational Unit not currently supported")
    }
}
// Check DC checks to ensure the domain component follows the path of worldpay -> local (will add other domain components)
func CheckDC(domain_component string) error {
    
    dc_name := strings.HasSuffix(domain_component,"=worldpay") || strings.HasSuffix(domain_component,"=local") 
    switch dc_name {
    case true:
        return nil
    default:
        return fmt.Errorf("Domain Component not currently supported")
    }
}
    
    
