/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    var l,r,ans=1,n,-1
    for ;l<=r;{
        var mid int = l+(r-l)/2;
        if isBadVersion(mid){
            ans=mid
            r=mid-1
        }else{
            l=(mid+1)
        }
    }
    return ans 
}