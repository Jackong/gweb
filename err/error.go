/**
 * User: Jackong
 * Date: 13-8-6
 * Time: 下午9:20
 */
package err

type Input string

func (this Input) Error() string {
	return string(this)
}
