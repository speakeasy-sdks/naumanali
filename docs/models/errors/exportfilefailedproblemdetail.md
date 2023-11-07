# ExportFileFailedProblemDetail

Represents when the export request is well-formed, but the server was unable to successfully export the requested file.


## Fields

| Field                                                                                                                                               | Type                                                                                                                                                | Required                                                                                                                                            | Description                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Detail`                                                                                                                                            | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A human-readable explanation specific to this occurrence of the problem.                                                                            |
| `ExportProblems`                                                                                                                                    | [][sdkerrors.ExportProblem](../../models/errors/exportproblem.md)                                                                                   | :heavy_check_mark:                                                                                                                                  | N/A                                                                                                                                                 |
| `Instance`                                                                                                                                          | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A URI reference that identifies the specific occurrence of the problem.                                                                             |
| `Status`                                                                                                                                            | [sdkerrors.SchemasExportFileFailedProblemDetailProblemDetailStatus](../../models/errors/schemasexportfilefailedproblemdetailproblemdetailstatus.md) | :heavy_check_mark:                                                                                                                                  | N/A                                                                                                                                                 |
| `Title`                                                                                                                                             | *string*                                                                                                                                            | :heavy_check_mark:                                                                                                                                  | A short, human-readable summary of the problem type.                                                                                                |
| `Type`                                                                                                                                              | *string*                                                                                                                                            | :heavy_check_mark:                                                                                                                                  | A URI reference that identifies the problem type.                                                                                                   |