package arm

import (
	"reflect"
	"testing"
)

var (
	dataSetID = "db01"
	basePath  = "../../testdata"

	HEADER = []string{
		"BegControl",
		"EndControl",
		"BegAtt",
		"EndAtt",
		"AccessControlListID_D",
		"Accuracy_1000063",
		"Application_1000124",
		"ArtifactID",
		"AttachmentDocumentIDs",
		"BatesBeg",
		"BatesBegAttach",
		"BatesEnd",
		"BatesEndAttach",
		"BCCAddress",
		"BrandingFont_1000230",
		"CategoryBasic_1000066",
		"CategoryNative_1000071",
		"CCAddress",
		"CenterFooterType_1000132",
		"ChartOrientation_1000061",
		"ChartType_1000057",
		"ChatID",
		"Classification_1000235",
		"Clearindentations_1000091",
		"ConfidentialityEndorsement",
		"ConnectionDate",
		"ConnectionTime",
		"ControlNumberEndAttach",
		"Conversation_ID",
		"Custodian",
		"DateAppStart",
		"DateAppStartDateOnly",
		"DateReceived",
		"DateReceivedDateOnly",
		"DateSent",
		"DateSentDateOnly",
		"Deduplicationmethod_1000148",
		"Deduplicationmethod_1000161",
		"Deleted_1000224",
		"Delimiter_1000153",
		"Delimiter_1000167",
		"DeNISTMode_1000150",
		"DeNISTMode_1000163",
		"DigitsforDocumentNumbering_1000127",
		"DiscoverStatus_1000174",
		"Display_1000054",
		"DisplaySMTPaddresses_1000095",
		"DitheringAlgorithm_1000097",
		"DocumentSubject",
		"DocumentType_1000187",
		"DownloadimagesfromInternet_1000092",
		"eDataAdmin_1000185",
		"EmailFrom",
		"EmailOutput_1000210",
		"EmailTo",
		"Errorstatus_1000158",
		"Errortype_1000160",
		"ExcelHeaderFooterExtraction_1000151",
		"ExcelHeaderFooterExtraction_1000165",
		"ExcelTextExtractionMethod_1000152",
		"ExcelTextExtractionMethod_1000166",
		"FileExtension",
		"FileIcon",
		"FileName",
		"FileSize",
		"Filtertype_1000173",
		"FolderorChatName_1000229",
		"Formatting_1000078",
		"FromAddress",
		"Generateplaceholderforeachfile_1000079",
		"GROUP_ID",
		"GroupByDateType_1000055",
		"GroupByResultLimit_1000142",
		"GroupIdentifier",
		"HasImages_1000036",
		"HasNative",
		"ImageSize_1000069",
		"ImagingMethod_1000214",
		"ImagingMethod_1000122",
		"IncludeBorders_1000098",
		"Includecomments_1000099",
		"Includegridlines_1000089",
		"Includeheadersandfooters_1000088",
		"Includerowandcolumnheadings_1000087",
		"InventoryStatus_1000176",
		"JobType_1000215",
		"Language_1000062",
		"LeftFooterType_1000131",
		"MD5Hash",
		"MessageID",
		"MLBAdmin_1000186",
		"MLBOriginalRelativityFolderPath",
		"MLBProduction_1000222",
		"NativeFile",
		"NearHit_1000236",
		"NumberingType_1000169",
		"NumberofAttachments",
		"NumberofDigits_1000154",
		"NumberofDigits_1000168",
		"OCRAccuracy_1000149",
		"OCRAccuracy_1000164",
		"Option_1000123",
		"Orientation_1000075",
		"OriginatingImagingDocumentError",
		"OriginatingNumber",
		"PageOrder_1000073",
		"PageOrientation_1000119",
		"Papersizeorientation_1000072",
		"ParentArtifactID_D",
		"ParentChildNumbering_1000155",
		"ParentChildNumbering_1000170",
		"ParentDate",
		"ParentDateDateOnly",
		"ParentDocumentID",
		"Phase2nearhit_1000237",
		"Phase2Tier1NearHit_1000238",
		"Phase2Tier2NearHit_1000239",
		"PivotOnDateType_1000056",
		"PivotOnResultLimit_1000143",
		"PlaceholderImageFormat_1000208",
		"PowerPointTextExtractionMethod_1000157",
		"PowerPointTextExtractionMethod_1000172",
		"Printarea_1000081",
		"Processingphase_1000159",
		"PROD_NUM",
		"ProdAttachRange",
		"ProdChildIDs",
		"ProdParentID",
		"ProductionFilename",
		"ProductionNumbering_1000126",
		"ProductionText",
		"ProductionType_1000141",
		"PublishStatus_1000175",
		"ReadDate",
		"ReadStatus_1000226",
		"ReadTime",
		"ReceivedTime",
		"RecipientCount",
		"RelativityImageCount",
		"RelativityNativeTimeZoneOffset",
		"RelativityNativeType",
		"RemovenonbreakingspaceNBSPcodes_1000084",
		"RendercolorpagestoJPEG_1000076",
		"RenderwithWord_1000083",
		"Resizeimagestofitpage_1000093",
		"Resizetablestofitpage_1000094",
		"RightFooterType_1000133",
		"SentTime",
		"SHA1Hash",
		"Show_1000080",
		"Showerrormessagesonplaceholders_1000077",
		"Showmessagetypeinheader_1000096",
		"Showspeakernotes_1000082",
		"Showtrackchanges_1000090",
		"Showtrackchangesinspreadsheet_1000086",
		"SlideOrientation_1000120",
		"Snippet",
		"SortOn_1000058",
		"SortOrder_1000059",
		"SortOrder1ASCDESC_1000136",
		"SortOrder2ASCDESC_1000137",
		"SortOrder3ASCDESC_1000138",
		"SortOrder4ASCDESC_1000139",
		"SortOrder5ASCDESC_1000140",
		"Source_1000070",
		"SourceLocation",
		"SOURCEPARTY",
		"StampIdentifier_1000241",
		"StampLocation_1000242",
		"Status_1000213",
		"Status_1000125",
		"SupportedByViewer",
		"TerminatingNumber",
		"TextDirection_1000225",
		"TextTypeCode_1000227",
		"Thread_ID",
		"TimeAppStart",
		"TimeZone_1000197",
		"TimeZoneField",
		"Title",
		"ToAddress",
		"Toggles_1000060",
		"Type_1000145",
		"Type_1000118",
		"Type_1000039",
		"Type_1000146",
		"Unhidehiddenworksheets_1000085",
		"UseImagePlaceholder_1000144",
		"Whenextractingchildrendonotextract_1000147",
		"Whenextractingchildrendonotextract_1000162",
		"WordTextExtractionMethod_1000156",
		"WordTextExtractionMethod_1000171",
		"XMLSourceFilename",
		"XMLSourceID",
		"XMLSourcePath",
		"NativePath",
		"TextPath",
	}

	expectedCountMap = map[string]int{
		"AccessControlListID_D":                      9993,
		"Accuracy_1000063":                           0,
		"Application_1000124":                        0,
		"ArtifactID":                                 9993,
		"AttachmentDocumentIDs":                      0,
		"BCCAddress":                                 0,
		"BatesBeg":                                   0,
		"BatesBegAttach":                             0,
		"BatesEnd":                                   0,
		"BatesEndAttach":                             0,
		"BegAtt":                                     9993,
		"BegControl":                                 9993,
		"BrandingFont_1000230":                       0,
		"CCAddress":                                  0,
		"CategoryBasic_1000066":                      0,
		"CategoryNative_1000071":                     0,
		"CenterFooterType_1000132":                   0,
		"ChartOrientation_1000061":                   0,
		"ChartType_1000057":                          0,
		"ChatID":                                     9993,
		"Classification_1000235":                     0,
		"Clearindentations_1000091":                  0,
		"ConfidentialityEndorsement":                 0,
		"ConnectionDate":                             0,
		"ConnectionTime":                             0,
		"ControlNumberEndAttach":                     0,
		"Conversation_ID":                            9993,
		"Custodian":                                  9993,
		"DateAppStart":                               0,
		"DateAppStartDateOnly":                       0,
		"DateReceived":                               4939,
		"DateReceivedDateOnly":                       0,
		"DateSent":                                   5053,
		"DateSentDateOnly":                           0,
		"DeNISTMode_1000150":                         0,
		"DeNISTMode_1000163":                         0,
		"Deduplicationmethod_1000148":                0,
		"Deduplicationmethod_1000161":                0,
		"Deleted_1000224":                            9993,
		"Delimiter_1000153":                          0,
		"Delimiter_1000167":                          0,
		"DigitsforDocumentNumbering_1000127":         0,
		"DiscoverStatus_1000174":                     0,
		"DisplaySMTPaddresses_1000095":               0,
		"Display_1000054":                            0,
		"DitheringAlgorithm_1000097":                 0,
		"DocumentSubject":                            0,
		"DocumentType_1000187":                       9993,
		"DownloadimagesfromInternet_1000092":         0,
		"EmailFrom":                                  4939,
		"EmailOutput_1000210":                        0,
		"EmailTo":                                    5054,
		"EndAtt":                                     9993,
		"EndControl":                                 9993,
		"Errorstatus_1000158":                        0,
		"Errortype_1000160":                          0,
		"ExcelHeaderFooterExtraction_1000151":        0,
		"ExcelHeaderFooterExtraction_1000165":        0,
		"ExcelTextExtractionMethod_1000152":          0,
		"ExcelTextExtractionMethod_1000166":          0,
		"FileExtension":                              0,
		"FileIcon":                                   9993,
		"FileName":                                   0,
		"FileSize":                                   0,
		"Filtertype_1000173":                         0,
		"FolderorChatName_1000229":                   9993,
		"Formatting_1000078":                         0,
		"FromAddress":                                0,
		"GROUP_ID":                                   0,
		"Generateplaceholderforeachfile_1000079":     0,
		"GroupByDateType_1000055":                    0,
		"GroupByResultLimit_1000142":                 0,
		"GroupIdentifier":                            9993,
		"HasImages_1000036":                          9993,
		"HasNative":                                  9993,
		"ImageSize_1000069":                          0,
		"ImagingMethod_1000122":                      0,
		"ImagingMethod_1000214":                      0,
		"IncludeBorders_1000098":                     0,
		"Includecomments_1000099":                    0,
		"Includegridlines_1000089":                   0,
		"Includeheadersandfooters_1000088":           0,
		"Includerowandcolumnheadings_1000087":        0,
		"InventoryStatus_1000176":                    0,
		"JobType_1000215":                            0,
		"Language_1000062":                           0,
		"LeftFooterType_1000131":                     0,
		"MD5Hash":                                    0,
		"MLBAdmin_1000186":                           0,
		"MLBOriginalRelativityFolderPath":            9993,
		"MLBProduction_1000222":                      0,
		"MessageID":                                  0,
		"NativeFile":                                 0,
		"NativePath":                                 9993,
		"NearHit_1000236":                            0,
		"NumberingType_1000169":                      0,
		"NumberofAttachments":                        0,
		"NumberofDigits_1000154":                     0,
		"NumberofDigits_1000168":                     0,
		"OCRAccuracy_1000149":                        0,
		"OCRAccuracy_1000164":                        0,
		"Option_1000123":                             0,
		"Orientation_1000075":                        0,
		"OriginatingImagingDocumentError":            0,
		"OriginatingNumber":                          0,
		"PROD_NUM":                                   0,
		"PageOrder_1000073":                          0,
		"PageOrientation_1000119":                    0,
		"Papersizeorientation_1000072":               0,
		"ParentArtifactID_D":                         9993,
		"ParentChildNumbering_1000155":               0,
		"ParentChildNumbering_1000170":               0,
		"ParentDate":                                 9993,
		"ParentDateDateOnly":                         9993,
		"ParentDocumentID":                           0,
		"Phase2Tier1NearHit_1000238":                 0,
		"Phase2Tier2NearHit_1000239":                 0,
		"Phase2nearhit_1000237":                      0,
		"PivotOnDateType_1000056":                    0,
		"PivotOnResultLimit_1000143":                 0,
		"PlaceholderImageFormat_1000208":             0,
		"PowerPointTextExtractionMethod_1000157":     0,
		"PowerPointTextExtractionMethod_1000172":     0,
		"Printarea_1000081":                          0,
		"Processingphase_1000159":                    0,
		"ProdAttachRange":                            0,
		"ProdChildIDs":                               0,
		"ProdParentID":                               0,
		"ProductionFilename":                         0,
		"ProductionNumbering_1000126":                0,
		"ProductionText":                             0,
		"ProductionType_1000141":                     0,
		"PublishStatus_1000175":                      0,
		"ReadDate":                                   0,
		"ReadStatus_1000226":                         0,
		"ReadTime":                                   0,
		"ReceivedTime":                               0,
		"RecipientCount":                             0,
		"RelativityImageCount":                       0,
		"RelativityNativeTimeZoneOffset":             0,
		"RelativityNativeType":                       9993,
		"RemovenonbreakingspaceNBSPcodes_1000084":    0,
		"RendercolorpagestoJPEG_1000076":             0,
		"RenderwithWord_1000083":                     0,
		"Resizeimagestofitpage_1000093":              0,
		"Resizetablestofitpage_1000094":              0,
		"RightFooterType_1000133":                    0,
		"SHA1Hash":                                   0,
		"SOURCEPARTY":                                0,
		"SentTime":                                   0,
		"Show_1000080":                               0,
		"Showerrormessagesonplaceholders_1000077":    0,
		"Showmessagetypeinheader_1000096":            0,
		"Showspeakernotes_1000082":                   0,
		"Showtrackchanges_1000090":                   0,
		"Showtrackchangesinspreadsheet_1000086":      0,
		"SlideOrientation_1000120":                   0,
		"Snippet":                                    0,
		"SortOn_1000058":                             0,
		"SortOrder1ASCDESC_1000136":                  0,
		"SortOrder2ASCDESC_1000137":                  0,
		"SortOrder3ASCDESC_1000138":                  0,
		"SortOrder4ASCDESC_1000139":                  0,
		"SortOrder5ASCDESC_1000140":                  0,
		"SortOrder_1000059":                          0,
		"SourceLocation":                             9993,
		"Source_1000070":                             0,
		"StampIdentifier_1000241":                    0,
		"StampLocation_1000242":                      0,
		"Status_1000125":                             0,
		"Status_1000213":                             0,
		"SupportedByViewer":                          9993,
		"TerminatingNumber":                          0,
		"TextDirection_1000225":                      0,
		"TextPath":                                   9993,
		"TextTypeCode_1000227":                       0,
		"Thread_ID":                                  0,
		"TimeAppStart":                               0,
		"TimeZoneField":                              0,
		"TimeZone_1000197":                           0,
		"Title":                                      0,
		"ToAddress":                                  0,
		"Toggles_1000060":                            0,
		"Type_1000039":                               0,
		"Type_1000118":                               0,
		"Type_1000145":                               0,
		"Type_1000146":                               0,
		"Unhidehiddenworksheets_1000085":             0,
		"UseImagePlaceholder_1000144":                0,
		"Whenextractingchildrendonotextract_1000147": 0,
		"Whenextractingchildrendonotextract_1000162": 0,
		"WordTextExtractionMethod_1000156":           0,
		"WordTextExtractionMethod_1000171":           0,
		"XMLSourceFilename":                          9993,
		"XMLSourceID":                                9993,
		"XMLSourcePath":                              0,
		"eDataAdmin_1000185":                         5150,
	}
)

func Test_Volumn_GetHeader(t *testing.T) {
	// Create a new Vol
	vol := &Volumn{
		BasePath:  basePath,
		DataSetID: dataSetID,
		ID:        "VOL001",
		OptFile:   "VOL001.opt",
		DatFile:   "VOL001.dat",
	}

	header, err := vol.ReadDatHeader()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if reflect.DeepEqual(header, HEADER) != true {
		t.Errorf("expected %v, but got %v", HEADER, header)
	}

}

func Test_Volumn_DatColumnValuesCount(t *testing.T) {

	vol := &Volumn{
		BasePath:  basePath,
		DataSetID: dataSetID,
		ID:        "VOL001",
		OptFile:   "VOL001.opt",
		DatFile:   "VOL001.dat",
	}

	countMap, err := vol.CountColumnWithValue()

	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if reflect.DeepEqual(expectedCountMap, countMap) != true {
		t.Errorf("expected %v, but got %v", expectedCountMap, countMap)
	}

}