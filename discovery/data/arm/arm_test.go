package arm

import (
	"reflect"
	"testing"
)

var (
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
)

func TestArmDBBuilder_Build(t *testing.T) {
	// Create a new ArmDBBuilder
	expectArmDB := &ARM{
		ID:       "db01",
		basePath: "../../testdata",
		Vols: []Vol{
			{
				basePath: "../../testdata",
				armID:    "db01",
				ID:       "VOL001",
				OptFile:  "VOL001.opt",
				DatFile:  "VOL001.dat",
			},
			{
				basePath: "../../testdata",
				armID:    "db01",
				ID:       "VOL002",
				OptFile:  "VOL002.opt",
				DatFile:  "VOL002.dat",
			},
			{
				basePath: "../../testdata",
				armID:    "db01",
				ID:       "VOL003",
				OptFile:  "VOL003.opt",
				DatFile:  "VOL003.dat",
			},
		},
	}

	builder := NewArmDBBuilder()
	db := builder.WithBasePath("../../testdata").WithID("db01").Build()

	err := db.Initialize()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if reflect.DeepEqual(expectArmDB, db) != true {
		t.Errorf("expected %v, but got %v", expectArmDB, db)
	}
}

func TestVolGetHeader(t *testing.T) {
	// Create a new Vol
	vol := &Vol{
		basePath: "../../testdata",
		armID:    "db01",
		ID:       "VOL001",
		OptFile:  "VOL001.opt",
		DatFile:  "VOL001.dat",
	}

	header, err := vol.ReadDatHeader()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if reflect.DeepEqual(header, HEADER) != true {
		t.Errorf("expected %v, but got %v", HEADER, header)
	}

}

func TestExport(t *testing.T) {

	vol := &Vol{
		basePath: "../../testdata",
		armID:    "db01",
		ID:       "VOL001",
		OptFile:  "VOL001.opt",
		DatFile:  "VOL001.dat",
	}

	// Call the Export function
	rowsChan := vol.StreamDatToChannel()

	for row := range rowsChan {
		if len(row) != len(HEADER) {
			t.Errorf("expected row length %d, but got %d", len(HEADER), len(row))
			return
		}
	}

}
