# BadRequestProblemDetail

Represents when the server could not understand the request due to invalid syntax.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Detail`                                                                    | **string*                                                                   | :heavy_minus_sign:                                                          | A human-readable explanation specific to this occurrence of the problem.    |
| `Instance`                                                                  | **string*                                                                   | :heavy_minus_sign:                                                          | A URI reference that identifies the specific occurrence of the problem.     |
| `InvalidParameters`                                                         | [][sdkerrors.InvalidParameter](../../models/errors/invalidparameter.md)     | :heavy_check_mark:                                                          | N/A                                                                         |
| `Status`                                                                    | [sdkerrors.ProblemDetailStatus](../../models/errors/problemdetailstatus.md) | :heavy_check_mark:                                                          | N/A                                                                         |
| `Title`                                                                     | *string*                                                                    | :heavy_check_mark:                                                          | A short, human-readable summary of the problem type.                        |
| `Type`                                                                      | *string*                                                                    | :heavy_check_mark:                                                          | A URI reference that identifies the problem type.                           |