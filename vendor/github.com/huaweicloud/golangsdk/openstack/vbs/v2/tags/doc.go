package tags

/*
Example to Add tag a Policy
	createopts := tags.CreateOpts{Tag:tags.ActionTags{Key:"test",Value:"demo"}}
	create,err := tags.Create(client,"ed8b9f73-4415-494d-a54e-5f3373bc353d",createopts).Extract()

Example to Get tag of a Policy
   get,err := tags.Get(client,"ed8b9f73-4415-494d-a54e-5f3373bc353d").Extract()
	fmt.Println(get,err)

Example to delete a tag of a Policy
    delete := tags.Delete(client,"5b549fad-c4e5-4d7e-83b9-eea366f27017","ECS")
	fmt.Println(delete)

Example to Query policy details based on tags
   queryOpts := tags.QueryOpts{Action:"filter",NotAnyTags:[]tags.QueryTags{{Key:"newKey",Values:[]string{"volumeSphere"}}}}
   query,err :=tags.Query(client,queryOpts).ExtractResources()
   fmt.Println(query,err)

Example to perform batch tag actions on a backup policy
	actionopts := tags.BatchOpts{Action:"update",Tags:[]tags.ActionTags{{Key:"k22",Value:"v22"}}}
    action := tags.BatchAction(client,"ed8b9f73-4415-494d-a54e-5f3373bc353d",actionopts)
*/
